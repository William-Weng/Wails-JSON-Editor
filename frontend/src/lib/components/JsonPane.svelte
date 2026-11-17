<script lang="ts">
  import { JSONEditor, Mode, type Content } from 'svelte-jsoneditor';
  import PaneHeader from './PaneHeader.svelte';
  import FontControls from './FontControls.svelte';

  let {
    content,
    fontSize = 14,
    minFontSize = 10,
    maxFontSize = 32,
    defaultFontSize = 14,
    hidden = false,
    onchange,
    onfontsizechange
  }: {
    content: Content;
    fontSize?: number;
    minFontSize?: number;
    maxFontSize?: number;
    defaultFontSize?: number;
    hidden?: boolean;
    onchange?: (content: Content) => void;
    onfontsizechange?: (value: number) => void;
  } = $props();
</script>

<section
  class="pane json-pane"
  class:is-hidden={hidden}
  aria-label="JSON"
>
  <PaneHeader title="JSON">
    <FontControls
      value={fontSize}
      min={minFontSize}
      max={maxFontSize}
      defaultValue={defaultFontSize}
      onchange={onfontsizechange}
    />
  </PaneHeader>

  <div
    class="editor-container"
    style="--jse-font-size-mono: {fontSize}px"
  >
      <JSONEditor
        {content}
        mode={Mode.tree}
        mainMenuBar={false}
        navigationBar={false}
        statusBar={false}
        onChange={(updatedContent) => onchange?.(updatedContent)}
      />
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

  .editor-container {
    flex: 1;
    min-width: 0;
    min-height: 0;
    --jse-border: none;
  }
</style>
