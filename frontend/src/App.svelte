<script lang="ts">
    import { Dialogs } from "@wailsio/runtime";
    import { onMount } from "svelte";
    import type { update } from "../wailsjs/go/models";
    import type { Content } from "svelte-jsoneditor";
    import { formatJsonError } from "./lib/utils/jsonError";
    import CodePane from "./lib/components/CodePane.svelte";
    import JsonPane from "./lib/components/JsonPane.svelte";
    import StatusBar from "./lib/components/StatusBar.svelte";
    import * as UpdateService from "../../frontend/bindings/wails-json-editor/backend/updateservice";

    type ViewMode = "text" | "json" | "split";

    const initialJson = {
        name: "JSON Editor",
        version: "1.0.0",
        enabled: true,
        settings: {
            theme: "dark",
            fontSize: 14,
        },
    };

    const DEFAULT_FONT_SIZE = 14;
    const MIN_FONT_SIZE = 10;
    const MAX_FONT_SIZE = 32;

    let sourceText = $state(JSON.stringify(initialJson, null, 2));
    let editorContent = $state.raw<Content>({ json: initialJson });
    let errorMessage = $state("");

    let codeFontSize = $state(DEFAULT_FONT_SIZE);
    let jsonFontSize = $state(DEFAULT_FONT_SIZE);
    let viewMode = $state<ViewMode>("split");

    const viewOptions: { value: ViewMode; label: string }[] = [
        { value: "text", label: "Text" },
        { value: "json", label: "JSON" },
        { value: "split", label: "Text JSON" },
    ];

    const lineCount = $derived(sourceText.split("\n").length);

    onMount(() => {
        checkUpdate();
    });

    function getErrorMessage(error: unknown) {
        return formatJsonError(error, sourceText);
    }

    async function checkUpdate() {
        const info = await UpdateService.CheckUpdate();
        console.log(info);

        if (!info?.hasUpdate) { return; }
        if (!info?.latestVersion) { return; }

        const result = await Dialogs.Question({
            Title: "有新版本請下載更新",
            Message: `發現新版本 v${info.latestVersion}，目前版本為 v${info.currentVersion}。是否前往下載？`,
            Buttons: [
                { Label: "取消", IsDefault: false },
                { Label: "確定", IsDefault: true },
            ],
        });

        if (result === "確定") {
            await UpdateService.OpenReleasePage();
        }
    }

    function applySourceText(text: string) {
        sourceText = text;

        try {
            const parsed = JSON.parse(text);
            editorContent = { json: parsed };
            errorMessage = "";
        } catch (error) {
            errorMessage = getErrorMessage(error);
        }
    }

    function handleEditorChange(content: Content) {
        editorContent = content;

        try {
            if ("json" in content) {
                sourceText = JSON.stringify(content.json, null, 2) ?? "";
            } else {
                sourceText = content.text ?? "";
            }

            JSON.parse(sourceText);
            errorMessage = "";
        } catch (error) {
            errorMessage = getErrorMessage(error);
        }
    }

    function transformJson(indent: number) {
        try {
            const parsed = JSON.parse(sourceText);
            const transformed = JSON.stringify(parsed, null, indent);

            sourceText = transformed ?? "";
            editorContent = { json: parsed };
            errorMessage = "";
        } catch (error) {
            errorMessage = getErrorMessage(error);
        }
    }

    const statusActions = [
        { label: "Minify", onclick: () => transformJson(0) },
        { label: "Format", onclick: () => transformJson(2) },
    ];
</script>

<div class="workspace">
    <main class="panes" class:single-pane={viewMode !== "split"}>
        <CodePane
            value={sourceText}
            fontSize={codeFontSize}
            minFontSize={MIN_FONT_SIZE}
            maxFontSize={MAX_FONT_SIZE}
            defaultFontSize={DEFAULT_FONT_SIZE}
            hidden={viewMode === "json"}
            oninput={applySourceText}
            onfontsizechange={(value) => (codeFontSize = value)}
        />

        <JsonPane
            content={editorContent}
            fontSize={jsonFontSize}
            minFontSize={MIN_FONT_SIZE}
            maxFontSize={MAX_FONT_SIZE}
            defaultFontSize={DEFAULT_FONT_SIZE}
            hidden={viewMode === "text"}
            onchange={handleEditorChange}
            onfontsizechange={(value) => (jsonFontSize = value)}
        />
    </main>

    <StatusBar
        {lineCount}
        {errorMessage}
        {viewMode}
        {viewOptions}
        actions={statusActions}
        onchangeviewmode={(value) => (viewMode = value)}
    />
</div>

<style>
    .workspace {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100dvh;
        min-height: 0;
        overflow: hidden;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    }

    .panes {
        display: grid;
        flex: 1;
        grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
        min-height: 0;
    }

    .single-pane {
        grid-template-columns: minmax(0, 1fr);
    }

    .panes:not(.single-pane) :global(.json-pane) {
        border-left: 1px solid var(--border);
    }
</style>
