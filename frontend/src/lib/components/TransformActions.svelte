<script lang="ts">
  type Action = {
    label: string;
    onclick: () => void;
  };

  let { actions }: { actions: Action[] } = $props();
</script>

<div class="transform-actions">
  {#each actions as action (action.label)}
    <button type="button" class="action-button" onclick={action.onclick}>
      {action.label}
    </button>
  {/each}
</div>

<style>
  .transform-actions {
    display: flex;
    align-items: center;
    flex-shrink: 0;
    gap: 8px;
  }

  .action-button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 28px;
    padding: 0 14px;
    border: 1px solid #385361;
    border-radius: 5px;
    background: #173642;
    color: #d5e4eb;
    font: inherit;
    font-weight: 600;
    white-space: nowrap;
    cursor: pointer;
    user-select: none;
  }

  .action-button:hover {
    border-color: #597b8d;
    background: #204553;
  }

  .action-button:active {
    background: #112b35;
  }

  .action-button:focus-visible {
    outline: 2px solid #7ca7d7;
    outline-offset: 2px;
  }
</style>
