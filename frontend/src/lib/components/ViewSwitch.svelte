<script lang="ts">
  type ViewMode = 'text' | 'json' | 'split';

  let {
    value,
    options,
    onchange
  }: {
    value: ViewMode;
    options: { value: ViewMode; label: string }[];
    onchange: (value: ViewMode) => void;
  } = $props();
</script>

<div class="view-switch" role="group" aria-label="檢視模式">
  {#each options as option (option.value)}
    <button
      type="button"
      class:active={value === option.value}
      aria-pressed={value === option.value}
      onclick={() => onchange(option.value)}
    >
      {option.label}
    </button>
  {/each}
</div>

<style>
  .view-switch {
    display: inline-flex;
    align-items: center;
    flex-shrink: 0;
    gap: 2px;
    padding: 2px;
    border: 1px solid #385361;
    border-radius: 6px;
    background: #081b23;
  }

  button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 24px;
    padding: 0 10px;
    border: none;
    border-radius: 4px;
    background: transparent;
    color: #8ba1ac;
    font: inherit;
    font-size: 12px;
    font-weight: 500;
    white-space: nowrap;
    cursor: pointer;
    user-select: none;
  }

  button:hover:not(.active) {
    background: #173642;
    color: #d5e4eb;
  }

  .active {
    background: #2a4c5b;
    color: #fff;
    box-shadow: 0 1px 3px rgb(0 0 0 / 15%);
  }

  button:focus-visible {
    outline: 2px solid #7ca7d7;
    outline-offset: 2px;
  }
</style>
