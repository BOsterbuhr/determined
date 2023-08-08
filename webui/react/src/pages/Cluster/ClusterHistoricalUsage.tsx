import { Space } from 'antd';
import dayjs from 'dayjs';
import React, { useCallback, useMemo, useState } from 'react';

import Button from 'components/kit/Button';
import Section from 'components/Section';
import { SyncProvider } from 'components/UPlot/SyncProvider';
import { useLoadable } from 'hooks/useLoadable';
import { useSettings } from 'hooks/useSettings';
import { getResourceAllocationAggregated } from 'services/api';
import userStore from 'stores/users';
import handleError from 'utils/error';
import { Loadable, NotLoaded } from 'utils/loadable';
import { useObservable } from 'utils/observable';

import css from './ClusterHistoricalUsage.module.scss';
import settingsConfig, { GroupBy, Settings } from './ClusterHistoricalUsage.settings';
import ClusterHistoricalUsageChart from './ClusterHistoricalUsageChart';
import ClusterHistoricalUsageCsvModal, { CSVGroupBy } from './ClusterHistoricalUsageCsvModal';
import ClusterHistoricalUsageFilters, {
  ClusterHistoricalUsageFiltersInterface,
} from './ClusterHistoricalUsageFilters';
import { mapResourceAllocationApiToChartSeries } from './utils';

export const DEFAULT_RANGE_DAY = 14;
export const DEFAULT_RANGE_MONTH = 6;
export const MAX_RANGE_DAY = 31;
export const MAX_RANGE_MONTH = 36;

const ClusterHistoricalUsage: React.FC = () => {
  const [isCsvModalVisible, setIsCsvModalVisible] = useState<boolean>(false);
  const { settings, updateSettings } = useSettings<Settings>(settingsConfig);
  const loadableUsers = useObservable(userStore.getUsers());
  const users = Loadable.getOrElse([], loadableUsers);

  const filters = useMemo(() => {
    const filters: ClusterHistoricalUsageFiltersInterface = {
      afterDate: dayjs().subtract(1 + DEFAULT_RANGE_DAY, 'day'),
      beforeDate: dayjs().subtract(1, 'day'),
      groupBy: GroupBy.Day,
    };

    if (settings.after) {
      const after = dayjs(settings.after || '');
      if (after.isValid() && after.isBefore(dayjs())) filters.afterDate = after;
    }
    if (settings.before) {
      const before = dayjs(settings.before || '');
      if (before.isValid() && before.isBefore(dayjs())) filters.beforeDate = before;
    }
    if (settings.groupBy && Object.values(GroupBy).includes(settings.groupBy as GroupBy)) {
      filters.groupBy = settings.groupBy as GroupBy;
    }

    // Validate filter dates.
    const dateDiff = filters.beforeDate.diff(filters.afterDate, filters.groupBy);
    if (filters.groupBy === GroupBy.Day && (dateDiff >= MAX_RANGE_DAY || dateDiff < 1)) {
      filters.afterDate = filters.beforeDate.clone().subtract(MAX_RANGE_DAY - 1, 'day');
    }
    if (filters.groupBy === GroupBy.Month && (dateDiff >= MAX_RANGE_MONTH || dateDiff < 1)) {
      filters.afterDate = filters.beforeDate.clone().subtract(MAX_RANGE_MONTH - 1, 'month');
    }

    return filters;
  }, [settings]);

  const aggRes = useLoadable(async () => {
    try {
      return await getResourceAllocationAggregated({
        endDate: filters.beforeDate,
        period:
          filters.groupBy === GroupBy.Month
            ? 'RESOURCE_ALLOCATION_AGGREGATION_PERIOD_MONTHLY'
            : 'RESOURCE_ALLOCATION_AGGREGATION_PERIOD_DAILY',
        startDate: filters.afterDate,
      });
    } catch (e) {
      handleError(e);
      return NotLoaded;
    }
  }, [filters.afterDate, filters.beforeDate, filters.groupBy]);

  const handleFilterChange = useCallback(
    (newFilter: ClusterHistoricalUsageFiltersInterface) => {
      const dateFormat = 'YYYY-MM' + (newFilter.groupBy === GroupBy.Day ? '-DD' : '');
      updateSettings({
        after: newFilter.afterDate.format(dateFormat),
        before: newFilter.beforeDate.format(dateFormat),
        groupBy: newFilter.groupBy,
      });
    },
    [updateSettings],
  );

  /**
   * When grouped by month force csv modal to display start/end of month.
   */
  let csvAfterDate = filters.afterDate;
  let csvBeforeDate = filters.beforeDate;
  if (filters.groupBy === GroupBy.Month) {
    csvAfterDate = csvAfterDate.startOf('month');
    csvBeforeDate = csvBeforeDate.endOf('month');
    if (csvBeforeDate.isAfter(dayjs())) {
      csvBeforeDate = dayjs().startOf('day');
    }
  }

  const chartSeries = useMemo(() => {
    return Loadable.map(aggRes, (response) => {
      return mapResourceAllocationApiToChartSeries(
        response.resourceEntries,
        filters.groupBy,
        users,
      );
    });
  }, [aggRes, filters.groupBy, users]);

  return (
    <div className={css.base}>
      <SyncProvider>
        <Space align="end" className={css.filters}>
          <ClusterHistoricalUsageFilters value={filters} onChange={handleFilterChange} />
          <Button onClick={() => setIsCsvModalVisible(true)}>Download CSV</Button>
        </Space>
        <Section
          bodyBorder
          loading={Loadable.isLoading(chartSeries)}
          title="Compute Hours Allocated">
          {Loadable.match(chartSeries, {
            Loaded: (series) => (
              <ClusterHistoricalUsageChart
                groupBy={series.groupedBy}
                hoursByLabel={series.hoursTotal}
                time={series.time}
              />
            ),
            NotLoaded: () => null,
          })}
        </Section>
        <Section
          bodyBorder
          loading={Loadable.isLoading(Loadable.all([loadableUsers, chartSeries]))}
          title="Compute Hours by User">
          {Loadable.match(chartSeries, {
            Loaded: (series) => (
              <ClusterHistoricalUsageChart
                groupBy={series.groupedBy}
                hoursByLabel={series.hoursByUsername}
                hoursTotal={series?.hoursTotal?.total}
                time={series.time}
              />
            ),
            NotLoaded: () => null,
          })}
        </Section>
        <Section
          bodyBorder
          loading={Loadable.isLoading(chartSeries)}
          title="Compute Hours by Label">
          {Loadable.match(chartSeries, {
            Loaded: (series) => (
              <ClusterHistoricalUsageChart
                groupBy={series.groupedBy}
                hoursByLabel={series.hoursByExperimentLabel}
                hoursTotal={series?.hoursTotal?.total}
                time={series.time}
              />
            ),
            NotLoaded: () => null,
          })}
        </Section>
        <Section
          bodyBorder
          loading={Loadable.isLoading(chartSeries)}
          title="Compute Hours by Resource Pool">
          {Loadable.match(chartSeries, {
            Loaded: (series) => (
              <ClusterHistoricalUsageChart
                groupBy={series.groupedBy}
                hoursByLabel={series.hoursByResourcePool}
                hoursTotal={series?.hoursTotal?.total}
                time={series.time}
              />
            ),
            NotLoaded: () => null,
          })}
        </Section>
        {isCsvModalVisible && (
          <ClusterHistoricalUsageCsvModal
            afterDate={csvAfterDate}
            beforeDate={csvBeforeDate}
            groupBy={CSVGroupBy.Workloads}
            onVisibleChange={setIsCsvModalVisible}
          />
        )}
      </SyncProvider>
    </div>
  );
};

export default ClusterHistoricalUsage;
