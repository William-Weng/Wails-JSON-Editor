<script lang="ts">
  let {
    value,
    min,
    max,
    defaultValue,
    onchange
  }: {
    value: number;
    min: number;
    max: number;
    defaultValue: number;
    onchange?: (value: number) => void;
  } = $props();
</script>

<div class="font-controls">
  <button
    type="button"
    aria-label="縮小字級"
    disabled={value <= min}
    onclick={() => onchange(value - 1)}
  >
    -
  </button>

  <span class="font-size">{value}px</span>

  <button
    type="button"
    aria-label="放大字級"
    disabled={value >= max}
    onclick={() => onchange(value + 1)}
  >
    +
  </button>

  <button
    type="button"
    class="reset-button"
    onclick={() => onchange(defaultValue)}
  >
    重設
  </button>
</div>

<style>
  .font-controls {
    display: flex;
    align-items: center;
    flex-shrink: 0;
    gap: 8px;
  }

  .font-size {
    min-width: 38px;
    font-size: 13px;
    font-weight: 600;
    font-variant-numeric: tabular-nums;
    text-align: center;
  }

  button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    padding: 0;
    border: 1px solid #cbd6e3;
    border-radius: 5px;
    background: #fff;
    color: #3d4b5f;
    font: inherit;
    font-size: 16px;
    cursor: pointer;
    user-select: none;
  }

  .reset-button {
    width: auto;
    padding: 0 9px;
    font-size: 12px;
  }

  button:disabled {
    opacity: 0.4;
    cursor: default;
  }

  button:hover:not(:disabled) {
    background: #edf2f8;
  }

  button:active:not(:disabled) {
    background: #e2e9f2;
  }

  button:focus-visible {
    outline: 2px solid #7ca7d7;
    outline-offset: 2px;
  }
</style>
