<script lang="ts">
  import ViewSwitch from './ViewSwitch.svelte';
  import TransformActions from './TransformActions.svelte';

  type ViewMode = 'text' | 'json' | 'split';

  let {
    lineCount,
    errorMessage,
    viewMode,
    viewOptions,
    actions,
    onchangeviewmode
  }: {
    lineCount: number;
    errorMessage: string;
    viewMode: ViewMode;
    viewOptions: { value: ViewMode; label: string }[];
    actions: { label: string; onclick: () => void }[];
    onchangeviewmode: (value: ViewMode) => void;
  } = $props();
</script>

<footer class="statusbar">
  <div class="status-info">
    <span class="line-count">{lineCount} 行</span>

    <span
      class="validation-message"
      class:invalid={Boolean(errorMessage)}
      title={errorMessage || 'JSON'}
      role="status"
    >
      {errorMessage || 'JSON'}
    </span>
  </div>

  <div class="status-actions">
    <ViewSwitch
      value={viewMode}
      options={viewOptions}
      onchange={onchangeviewmode}
    />

    <TransformActions {actions} />
  </div>
</footer>

<style>
  .statusbar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex: 0 0 40px;
    min-width: 0;
    gap: 16px;
    padding: 0 12px 0 16px;
    border-top: 1px solid #263d47;
    background: var(--status-bg);
    color: var(--status-text);
    font-size: 12px;
  }

  .status-info {
    display: flex;
    align-items: center;
    flex: 1;
    min-width: 0;
    gap: 16px;
  }

  .line-count {
    flex-shrink: 0;
    white-space: nowrap;
  }

  .validation-message {
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .invalid {
    color: var(--error-text);
  }

  .status-actions {
    display: flex;
    align-items: center;
    flex-shrink: 0;
    gap: 8px;
  }
</style>
