import { SortOrder } from 'antd/es/table/interface';
import { debounce } from 'lodash';
import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';

import dropdownCss from 'components/ActionDropdown/ActionDropdown.module.scss';
import AddUsersToGroupsModalComponent from 'components/AddUsersToGroupsModal';
import ChangeUserStatusModalComponent from 'components/ChangeUserStatusModal';
import ConfigureAgentModalComponent from 'components/ConfigureAgentModal';
import CreateUserModalComponent from 'components/CreateUserModal';
import Button from 'components/kit/Button';
import { Column, Columns } from 'components/kit/Columns';
import Dropdown, { MenuItem } from 'components/kit/Dropdown';
import Icon from 'components/kit/Icon';
import Input from 'components/kit/Input';
import { useModal } from 'components/kit/Modal';
import Select, { SelectValue } from 'components/kit/Select';
import { makeToast } from 'components/kit/Toast';
import { Loadable, NotLoaded } from 'components/kit/utils/loadable';
import ManageGroupsModalComponent from 'components/ManageGroupsModal';
import Section from 'components/Section';
import SetUserRolesModalComponent from 'components/SetUserRolesModal';
import InteractiveTable, { onRightClickableCell } from 'components/Table/InteractiveTable';
import SkeletonTable from 'components/Table/SkeletonTable';
import {
  checkmarkRenderer,
  defaultRowClassName,
  getFullPaginationConfig,
  relativeTimeRenderer,
} from 'components/Table/Table';
import UserBadge from 'components/UserBadge';
import { useLoadable } from 'hooks/useLoadable';
import usePermissions from 'hooks/usePermissions';
import { getGroups, getUsers, patchUsers } from 'services/api';
import { V1GetUsersRequestSortBy, V1GroupSearchResult, V1OrderBy } from 'services/api-ts-sdk';
import determinedStore from 'stores/determinedInfo';
import roleStore from 'stores/roles';
import userSettings from 'stores/userSettings';
import { DetailedUser } from 'types';
import handleError, { ErrorType } from 'utils/error';
import { useObservable } from 'utils/observable';
import { validateDetApiEnum } from 'utils/service';
import { alphaNumericSorter, booleanSorter, numericSorter } from 'utils/sort';

import css from './UserManagement.module.scss';
import {
  DEFAULT_COLUMN_WIDTHS,
  DEFAULT_SETTINGS,
  UserManagementSettings,
  UserRole,
  UserStatus,
} from './UserManagement.settings';

export const USER_TITLE = 'Users';
export const CREATE_USER = 'Add User';
export const CREATE_USER_LABEL = 'add_user';

interface DropdownProps {
  fetchUsers: () => void;
  groups: V1GroupSearchResult[];
  user: DetailedUser;
  userManagementEnabled: boolean;
}

const MenuKey = {
  Agent: 'agent',
  Edit: 'edit',
  Groups: 'groups',
  State: 'state',
  View: 'view',
} as const;

const ActionMenuKey = {
  AddToGroups: 'add-to-groups',
  ChangeStatus: 'change-status',
  SetRoles: 'set-roles',
} as const;

const UserActionDropdown = ({ fetchUsers, user, groups, userManagementEnabled }: DropdownProps) => {
  const EditUserModal = useModal(CreateUserModalComponent);
  const ViewUserModal = useModal(CreateUserModalComponent);
  const ManageGroupsModal = useModal(ManageGroupsModalComponent);
  const ConfigureAgentModal = useModal(ConfigureAgentModalComponent);
  const [selectedUserGroups, setSelectedUserGroups] = useState<V1GroupSearchResult[]>();

  const { canModifyUsers } = usePermissions();
  const { rbacEnabled } = useObservable(determinedStore.info);

  const onToggleActive = useCallback(async () => {
    try {
      await patchUsers({ activate: !user.isActive, userIds: [user.id] });
      makeToast({
        severity: 'Confirm',
        title: `User has been ${user.isActive ? 'deactivated' : 'activated'}`,
      });
      fetchUsers();
    } catch (e) {
      handleError(e, {
        isUserTriggered: true,
        publicSubject: `Unable to ${user.isActive ? 'deactivate' : 'activate'} user.`,
        silent: false,
        type: ErrorType.Api,
      });
    }
  }, [fetchUsers, user]);

  const menuItems =
    userManagementEnabled && canModifyUsers
      ? rbacEnabled
        ? [
            { key: MenuKey.Edit, label: 'Edit User' },
            { key: MenuKey.Groups, label: 'Manage Groups' },
            { key: MenuKey.Agent, label: 'Configure Agent' },
            { key: MenuKey.State, label: `${user.isActive ? 'Deactivate' : 'Activate'}` },
          ]
        : [
            { key: MenuKey.Edit, label: 'Edit User' },
            { key: MenuKey.Agent, label: 'Configure Agent' },
            { key: MenuKey.State, label: `${user.isActive ? 'Deactivate' : 'Activate'}` },
          ]
      : [{ key: MenuKey.View, label: 'View User' }];

  const handleDropdown = useCallback(
    async (key: string) => {
      switch (key) {
        case MenuKey.Agent:
          ConfigureAgentModal.open();
          break;
        case MenuKey.Edit:
          EditUserModal.open();
          break;
        case MenuKey.Groups: {
          const response = await getGroups({ limit: 500, userId: user.id });
          setSelectedUserGroups(response.groups ?? []);
          ManageGroupsModal.open();
          break;
        }
        case MenuKey.State:
          await onToggleActive();
          break;
        case MenuKey.View:
          ViewUserModal.open();
          break;
      }
    },
    [ConfigureAgentModal, EditUserModal, ManageGroupsModal, onToggleActive, user, ViewUserModal],
  );

  return (
    <div className={dropdownCss.base}>
      <Dropdown menu={menuItems} placement="bottomRight" onClick={handleDropdown}>
        <Button icon={<Icon name="overflow-vertical" size="small" title="Action menu" />} />
      </Dropdown>
      <ViewUserModal.Component user={user} viewOnly onClose={fetchUsers} />
      <EditUserModal.Component user={user} onClose={fetchUsers} />
      <ManageGroupsModal.Component
        groupOptions={groups}
        user={user}
        userGroups={selectedUserGroups ?? []}
      />
      <ConfigureAgentModal.Component user={user} onClose={fetchUsers} />
    </div>
  );
};

const roleOptions = [
  { label: 'Admin', value: UserRole.ADMIN },
  { label: 'Member', value: UserRole.MEMBER },
];

const statusOptions = [
  { label: 'Active', value: UserStatus.ACTIVE },
  { label: 'Inactive', value: UserStatus.INACTIVE },
];

const userManagementSettings = userSettings.get(UserManagementSettings, 'user-management');
const UserManagement: React.FC = () => {
  const [selectedUserIds, setSelectedUserIds] = useState<React.Key[]>([]);
  const [refresh, setRefresh] = useState({});
  const pageRef = useRef<HTMLElement>(null);
  const loadableSettings = useObservable(userManagementSettings);
  const settings = useMemo(() => {
    return Loadable.match(loadableSettings, {
      _: () => DEFAULT_SETTINGS,
      Loaded: (s) => ({ ...DEFAULT_SETTINGS, ...s }),
    });
  }, [loadableSettings]);
  const updateSettings = useCallback(
    (p: Partial<UserManagementSettings>) =>
      userSettings.setPartial(UserManagementSettings, 'user-management', p),
    [],
  );

  const userResponse = useLoadable(async () => {
    try {
      return await getUsers({
        active: settings.statusFilter && settings.statusFilter === UserStatus.ACTIVE,
        admin: settings.roleFilter && settings.roleFilter === UserRole.ADMIN,
        limit: settings.tableLimit,
        name: settings.name,
        offset: settings.tableOffset,
        orderBy: settings.sortDesc ? V1OrderBy.DESC : V1OrderBy.ASC,
        sortBy: settings.sortKey || V1GetUsersRequestSortBy.UNSPECIFIED,
      });
    } catch (e) {
      handleError(e, { publicSubject: 'Could not fetch user search results' });
      return NotLoaded;
    }
  }, [settings, refresh]);

  const users = useMemo(
    () =>
      Loadable.match(userResponse, {
        Loaded: (r) => r.users,
        _: () => [],
      }),
    [userResponse],
  );

  const { rbacEnabled } = useObservable(determinedStore.info);
  const { canModifyUsers, canModifyPermissions } = usePermissions();
  const info = useObservable(determinedStore.info);
  const ChangeUserStatusModal = useModal(ChangeUserStatusModalComponent);
  const SetUserRolesModal = useModal(SetUserRolesModalComponent);
  const AddUsersToGroupsModal = useModal(AddUsersToGroupsModalComponent);

  const fetchUsers = useCallback((): void => {
    if (!settings) return;

    setRefresh({});
  }, [settings, setRefresh]);

  const groupsResponse = useLoadable(async (canceler) => {
    try {
      return await getGroups({ limit: 500 }, { signal: canceler.signal });
    } catch (e) {
      handleError(e, { publicSubject: 'Unable to fetch groups.' });
      return NotLoaded;
    }
  }, []);
  const groups = useMemo(
    () =>
      Loadable.match(groupsResponse, {
        Loaded: (g) => g.groups || [],
        _: () => [],
      }),
    [groupsResponse],
  );

  useEffect(() => (rbacEnabled ? roleStore.fetch() : undefined), [rbacEnabled]);

  useEffect(() => {
    // reset invalid settings
    if (Loadable.isLoaded(loadableSettings) && !Loadable.getOrElse(null, loadableSettings)) {
      updateSettings(DEFAULT_SETTINGS);
    }
  }, [loadableSettings, updateSettings]);

  const CreateUserModal = useModal(CreateUserModalComponent);

  const handleNameSearchApply = useMemo(
    () =>
      debounce(
        (e: React.ChangeEvent<HTMLInputElement>) =>
          updateSettings({ name: e.target.value || undefined, row: undefined, tableOffset: 0 }),
        500,
      ),
    [updateSettings],
  );

  const handleStatusFilterApply = useCallback(
    (statusFilter?: SelectValue) =>
      updateSettings({ row: undefined, statusFilter: statusFilter as UserStatus, tableOffset: 0 }),
    [updateSettings],
  );

  const handleRoleFilterApply = useCallback(
    (roleFilter?: SelectValue) =>
      updateSettings({ roleFilter: roleFilter as UserRole, row: undefined, tableOffset: 0 }),
    [updateSettings],
  );

  const handleTableRowSelect = useCallback((rowKeys: React.Key[]) => {
    setSelectedUserIds(rowKeys);
  }, []);

  const actionDropdownMenu: MenuItem[] = useMemo(() => {
    const menuItems: MenuItem[] = [{ key: ActionMenuKey.ChangeStatus, label: 'Change Status' }];

    if (rbacEnabled) {
      if (canModifyPermissions) {
        menuItems.push({ key: ActionMenuKey.SetRoles, label: 'Set Roles' });
      }
      if (canModifyUsers) {
        menuItems.push({ key: ActionMenuKey.AddToGroups, label: 'Add to Groups' });
      }
    }

    return menuItems;
  }, [rbacEnabled, canModifyPermissions, canModifyUsers]);

  const handleActionDropdown = useCallback(
    (key: string) => {
      switch (key) {
        case ActionMenuKey.AddToGroups:
          AddUsersToGroupsModal.open();
          break;
        case ActionMenuKey.ChangeStatus:
          ChangeUserStatusModal.open();
          break;
        case ActionMenuKey.SetRoles:
          SetUserRolesModal.open();
          break;
      }
    },
    [AddUsersToGroupsModal, ChangeUserStatusModal, SetUserRolesModal],
  );

  const clearTableSelection = useCallback(() => {
    setSelectedUserIds([]);
  }, []);

  const filterIcon = useCallback(() => <Icon name="search" size="tiny" title="Search" />, []);

  const columns = useMemo(() => {
    const actionRenderer = (_: string, record: DetailedUser) => {
      return (
        <UserActionDropdown
          fetchUsers={fetchUsers}
          groups={groups}
          user={record}
          userManagementEnabled={info.userManagementEnabled}
        />
      );
    };
    const defaultSortKey: V1GetUsersRequestSortBy = validateDetApiEnum(
      V1GetUsersRequestSortBy,
      settings.sortKey,
    );
    const defaultSortOrder: SortOrder = settings.sortDesc ? 'descend' : 'ascend';
    const columns = [
      {
        dataIndex: 'displayName',
        defaultSortOrder:
          defaultSortKey === V1GetUsersRequestSortBy.NAME ? defaultSortOrder : undefined,
        defaultWidth: DEFAULT_COLUMN_WIDTHS['displayName'],
        key: V1GetUsersRequestSortBy.NAME,
        onCell: onRightClickableCell,
        render: (_: string, r: DetailedUser) => <UserBadge user={r} />,
        sorter: (a: DetailedUser, b: DetailedUser) => {
          return alphaNumericSorter(a.displayName || a.username, b.displayName || b.username);
        },
        title: 'Name',
      },
      {
        dataIndex: 'isActive',
        defaultSortOrder:
          defaultSortKey === V1GetUsersRequestSortBy.ACTIVE ? defaultSortOrder : undefined,
        defaultWidth: DEFAULT_COLUMN_WIDTHS['isActive'],
        key: V1GetUsersRequestSortBy.ACTIVE,
        onCell: onRightClickableCell,
        render: checkmarkRenderer,
        sorter: (a: DetailedUser, b: DetailedUser) => booleanSorter(a.isActive, b.isActive),
        title: 'Active',
      },
      {
        dataIndex: 'isAdmin',
        defaultSortOrder:
          defaultSortKey === V1GetUsersRequestSortBy.ADMIN ? defaultSortOrder : undefined,
        defaultWidth: DEFAULT_COLUMN_WIDTHS['isAdmin'],
        key: V1GetUsersRequestSortBy.ADMIN,
        onCell: onRightClickableCell,
        render: checkmarkRenderer,
        sorter: (a: DetailedUser, b: DetailedUser) => booleanSorter(a.isAdmin, b.isAdmin),
        title: 'Admin',
      },
      {
        dataIndex: 'modifiedAt',
        defaultSortOrder:
          defaultSortKey === V1GetUsersRequestSortBy.MODIFIEDTIME ? defaultSortOrder : undefined,
        defaultWidth: DEFAULT_COLUMN_WIDTHS['modifiedAt'],
        key: V1GetUsersRequestSortBy.MODIFIEDTIME,
        onCell: onRightClickableCell,
        render: (value: number): React.ReactNode => relativeTimeRenderer(new Date(value)),
        sorter: (a: DetailedUser, b: DetailedUser) => numericSorter(a.modifiedAt, b.modifiedAt),
        title: 'Modified Time',
      },
      {
        dataIndex: 'lastAuthAt',
        defaultSortOrder:
          defaultSortKey === V1GetUsersRequestSortBy.LASTAUTHTIME ? defaultSortOrder : undefined,
        defaultWidth: DEFAULT_COLUMN_WIDTHS['lastAuthAt'],
        key: V1GetUsersRequestSortBy.LASTAUTHTIME,
        onCell: onRightClickableCell,
        render: (value: number | undefined): React.ReactNode => {
          return value ? (
            relativeTimeRenderer(new Date(value))
          ) : (
            <div className={css.rightAligned}>N/A</div>
          );
        },
        sorter: (a: DetailedUser, b: DetailedUser) => numericSorter(a.lastAuthAt, b.lastAuthAt),
        title: 'Last Seen',
      },
      {
        className: 'fullCell',
        dataIndex: 'action',
        defaultWidth: DEFAULT_COLUMN_WIDTHS['action'],
        key: 'action',
        onCell: onRightClickableCell,
        render: actionRenderer,
        title: '',
        width: DEFAULT_COLUMN_WIDTHS['action'],
      },
    ];
    return rbacEnabled ? columns.filter((c) => c.dataIndex !== 'isAdmin') : columns;
  }, [
    fetchUsers,
    filterIcon,
    groups,
    info.userManagementEnabled,
    rbacEnabled,
    settings,
  ]);

  return (
    <>
      <Section className={css.usersTable}>
        <Columns>
          <Column>
            <Columns>
              {/* input is uncontrolled to prevent settings overwriting during composition */}
              <Input
                defaultValue={settings.name}
                prefix={<Icon color="cancel" decorative name="search" />}
                onChange={handleNameSearchApply}
              />
              <Select
                allowClear
                options={roleOptions}
                placeholder="All roles"
                searchable={false}
                value={settings.roleFilter}
                onChange={handleRoleFilterApply}
              />
              <Select
                allowClear
                options={statusOptions}
                placeholder="All statuses"
                searchable={false}
                value={settings.statusFilter}
                onChange={handleStatusFilterApply}
              />
            </Columns>
          </Column>
          <Column align="right">
            {selectedUserIds.length > 0 && (
              <Dropdown menu={actionDropdownMenu} onClick={handleActionDropdown}>
                <Button>Actions</Button>
              </Dropdown>
            )}
            <Button
              aria-label={CREATE_USER_LABEL}
              disabled={!info.userManagementEnabled || !canModifyUsers}
              onClick={CreateUserModal.open}>
              {CREATE_USER}
            </Button>
          </Column>
        </Columns>
        {settings ? (
          <InteractiveTable<DetailedUser, UserManagementSettings>
            columns={columns}
            containerRef={pageRef}
            dataSource={users}
            interactiveColumns={false}
            loading={Loadable.isNotLoaded(userResponse)}
            pagination={Loadable.match(userResponse, {
              _: () => undefined,
              Loaded: (r) =>
                getFullPaginationConfig(
                  {
                    limit: settings.tableLimit,
                    offset: settings.tableOffset,
                  },
                  r.pagination.total || 0,
                ),
            })}
            rowClassName={defaultRowClassName({ clickable: false })}
            rowKey="id"
            settings={settings}
            showSorterTooltip={false}
            size="small"
            updateSettings={updateSettings}
          />
        ) : (
          <SkeletonTable columns={columns.length} />
        )}
      </Section>
      <CreateUserModal.Component onClose={fetchUsers} />
      <ChangeUserStatusModal.Component
        clearTableSelection={clearTableSelection}
        fetchUsers={fetchUsers}
        userIds={selectedUserIds.map((id) => Number(id))}
      />
      <SetUserRolesModal.Component
        clearTableSelection={clearTableSelection}
        fetchUsers={fetchUsers}
        userIds={selectedUserIds.map((id) => Number(id))}
      />
      <AddUsersToGroupsModal.Component
        clearTableSelection={clearTableSelection}
        fetchUsers={fetchUsers}
        groupOptions={groups}
        userIds={selectedUserIds.map((id) => Number(id))}
      />
    </>
  );
};

export default UserManagement;
