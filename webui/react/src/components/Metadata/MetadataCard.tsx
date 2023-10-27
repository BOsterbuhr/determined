import { Card, Space } from 'antd';
import Button from 'determined-ui/Button';
import Icon from 'determined-ui/Icon';
import Spinner from 'determined-ui/Spinner';
import React, { useCallback, useMemo, useRef, useState } from 'react';

import { Metadata } from 'types';
import handleError, { ErrorType } from 'utils/error';

import EditableMetadata from './EditableMetadata';

interface Props {
  disabled?: boolean;
  metadata?: Metadata;
  onSave?: (newMetadata: Metadata) => Promise<void>;
}

const MetadataCard: React.FC<Props> = ({ disabled = false, metadata = {}, onSave }: Props) => {
  const containerRef = useRef(null);
  const [isEditing, setIsEditing] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [editedMetadata, setEditedMetadata] = useState<Metadata>(metadata ?? {});

  const metadataArray = useMemo(() => {
    return Object.entries(metadata ?? {}).map(([key, value]) => {
      return { content: value, label: key };
    });
  }, [metadata]);

  const editMetadata = useCallback(() => {
    if (disabled) return;
    setIsEditing(true);
  }, [disabled]);

  const saveMetadata = useCallback(async () => {
    try {
      setIsLoading(true);
      await onSave?.(editedMetadata);
      setIsEditing(false);
    } catch (e) {
      handleError(containerRef, e, {
        publicSubject: 'Unable to update metadata.',
        silent: true,
        type: ErrorType.Api,
      });
    }
    setIsLoading(false);
  }, [editedMetadata, onSave]);

  const cancelEditMetadata = useCallback(() => {
    setEditedMetadata(metadata);
    setIsEditing(false);
  }, [setEditedMetadata, metadata]);

  const showPlaceholder = useMemo(() => {
    return metadataArray.length === 0 && !isEditing;
  }, [isEditing, metadataArray.length]);

  return (
    <div ref={containerRef}>
      <Card
        bodyStyle={{ padding: '16px' }}
        extra={
          isEditing ? (
            <Space size="small">
              <Button size="small" onClick={cancelEditMetadata}>
                Cancel
              </Button>
              <Button size="small" type="primary" onClick={saveMetadata}>
                Save
              </Button>
            </Space>
          ) : (
            disabled || (
              <Button
                icon={<Icon name="pencil" showTooltip size="small" title="Edit" />}
                type="text"
                onClick={editMetadata}
              />
            )
          )
        }
        headStyle={{ paddingInline: '16px' }}
        title={'Metadata'}>
        {showPlaceholder ? (
          <div
            style={{ color: 'var(--theme-colors-monochrome-9)', fontStyle: 'italic' }}
            onClick={editMetadata}>
            {disabled ? 'No metadata present.' : 'Add Metadata...'}
          </div>
        ) : (
          <Spinner spinning={isLoading}>
            <EditableMetadata
              editing={isEditing}
              metadata={editedMetadata}
              updateMetadata={setEditedMetadata}
            />
          </Spinner>
        )}
      </Card>
    </div>
  );
};

export default MetadataCard;
