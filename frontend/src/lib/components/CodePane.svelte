<script lang="ts">
  import PaneHeader from './PaneHeader.svelte';
  import FontControls from './FontControls.svelte';

  let {
    value = '',
    fontSize = 14,
    minFontSize = 10,
    maxFontSize = 32,
    defaultFontSize = 14,
    hidden = false,
    oninput,
    onfontsizechange
  }: {
    value?: string;
    fontSize?: number;
    minFontSize?: number;
    maxFontSize?: number;
    defaultFontSize?: number;
    hidden?: boolean;
    oninput?: (value: string) => void;
    onfontsizechange?: (value: number) => void;
  } = $props();

  let textarea: HTMLTextAreaElement | undefined = $state();
  let lineNumbers: HTMLDivElement | undefined = $state();

  const lineCount = $derived(value.split('\n').length);

  function handleInput(event: Event) {
    const target = event.currentTarget as HTMLTextAreaElement;
    oninput?.(target.value);
  }

  function syncScroll() {
    if (!textarea || !lineNumbers) return;
    lineNumbers.scrollTop = textarea.scrollTop;
  }
</script>

<section
  class="pane code-pane"
  class:is-hidden={hidden}
  aria-label="JSON"
>
  <PaneHeader title="CODE">
    <FontControls
      value={fontSize}
      min={minFontSize}
      max={maxFontSize}
      defaultValue={defaultFontSize}
      onchange={onfontsizechange}
    />
  </PaneHeader>

  <div
    class="text-editor"
    style="--code-font-size: {fontSize}px; --code-line-height: {fontSize * 1.8}px"
  >
    <div class="line-numbers" bind:this={lineNumbers} aria-hidden="true">
      {#each Array(lineCount) as _, index}
        <div class="line-number">{index + 1}</div>
      {/each}
    </div>

    <textarea
      bind:this={textarea}
      aria-label="JSON"
      value={value}
      oninput={handleInput}
      onscroll={syncScroll}
      spellcheck="false"
      autocapitalize="off"
      autocomplete="off"
      wrap="off"
    ></textarea>
  </div>
</section>

<style>
  .pane {
    display: flex;
    flex-direction: column;
    min-width: 0;
    min-height: 0;
    overflow: hidden;
  }

  .is-hidden {
    display: none;
  }

  .text-editor {
    display: flex;
    flex: 1;
    min-width: 0;
    min-height: 0;
    overflow: hidden;
    background: var(--editor-bg);
  }

  .line-numbers,
  textarea {
    font-family: SFMono-Regular, Menlo, Monaco, Consolas, monospace;
    font-size: var(--code-font-size);
    line-height: var(--code-line-height);
    font-variant-ligatures: none;
  }

  .line-numbers {
    flex: 0 0 auto;
    min-width: 46px;
    height: 100%;
    padding: 16px 10px;
    overflow: hidden;
    border-right: 1px solid #203b46;
    background: var(--gutter-bg);
    color: var(--gutter-text);
    text-align: right;
    user-select: none;
  }

  .line-number {
    height: var(--code-line-height);
  }

  textarea {
    display: block;
    flex: 1;
    width: 0;
    min-width: 0;
    height: 100%;
    margin: 0;
    padding: 16px;
    resize: none;
    overflow: auto;
    border: none;
    border-radius: 0;
    outline: none;
    background: transparent;
    color: var(--editor-text);
    caret-color: #fff;
    tab-size: 2;
    white-space: pre;
  }

  textarea::selection {
    background: #34596b;
  }
</style>
