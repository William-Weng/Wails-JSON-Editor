<script lang="ts">
  import type { Snippet } from 'svelte';

  interface Props {
    title: string;
    children?: Snippet;
  }

  let { title, children }: Props = $props();
</script>

<header class="pane-header">
  <span class="pane-title">{title}</span>

  <div class="pane-header-actions">
    {@render children?.()}
  </div>
</header>

<style>
  .pane-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex: 0 0 40px;
    gap: 12px;
    padding: 0 16px;
    border-bottom: 1px solid var(--border, #d6dfe9);
    background: var(--header-bg, #f1f4f9);
    color: var(--header-text, #657790);
    user-select: none;
  }

  .pane-title {
    font-size: 13px;
    font-weight: 700;
    letter-spacing: 1px;
  }

  .pane-header-actions {
    display: flex;
    align-items: center;
    flex-shrink: 0;
    gap: 8px;
  }
</style>
