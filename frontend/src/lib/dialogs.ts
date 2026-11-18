import { Dialogs } from "@wailsio/runtime";

/**
 * QuestionDialog 的選項
 */
interface QuestionDialogOptions {
  title: string; // 對話框標題
  message: string; // 對話框內文，可使用 \n 換行
  buttons: {
    // 按鈕清單；順序會影響 dialog 中按鈕的顯示順序
    label: string; // 按鈕文字，也是 Dialogs.Question() 回傳的結果值
    isDefault?: boolean; // 是否為預設按鈕；通常對應鍵盤 Enter
  }[];
}

/**
 * 顯示只有一顆按鈕的提示對話框
 *
 * 適合用於：
 * - 操作成功或失敗的通知
 * - 錯誤訊息
 * - 「已是最新版本」
 * - 「稍後再檢查更新」等單一選擇情境
 *
 * @param options - 對話框標題、訊息與按鈕文字
 * @returns 使用者點擊的按鈕文字（Label）
 *
 * @example
 * await showAlertDialog({
 *   title: "更新檢查失敗",
 *   message: "無法連線到 GitHub，請稍後再試。",
 *   buttonLabel: "好",
 * });
 *
 * @example
 * await showAlertDialog({
 *   title: "已是最新版本",
 *   message: "目前已是最新版本。",
 * });
 */
export async function showAlertDialog({
	title,
	message,
	buttonLabel = "好",
}: {
	title: string;
	message: string;
	buttonLabel?: string;
}): Promise<string> {
	return await showQuestionDialog({
		title,
		message,
		buttons: [
			{
				label: buttonLabel,
				isDefault: true,
			},
		],
	});
}

/**
 * 顯示常見的「取消 / 確定」對話框
 *
 * 這是 showQuestionDialog() 的便捷版本，合只需要二選一的情境，例如刪除確認、更新確認等
 *
 * @param options - 對話框標題、訊息，以及可自訂的按鈕文字
 * @returns 使用者是否按下確認按鈕
 *
 * @example
 * const confirmed = await showConfirmDialog({
 *   title: "刪除確認",
 *   message: "確定要刪除嗎？",
 *   confirmLabel: "刪除",
 * });
 *
 * if (confirmed) {
 *   await deleteItem();
 * }
 */
export async function showConfirmDialog({
  title,
  message,
  confirmLabel = "確定",
  cancelLabel = "取消",
}: {
  title: string;
  message: string;
  confirmLabel?: string;
  cancelLabel?: string;
}): Promise<boolean> {
  const result = await showQuestionDialog({
    title,
    message,
    buttons: [
      {
        label: cancelLabel,
        isDefault: false,
      },
      {
        label: confirmLabel,
        isDefault: true,
      },
    ],
  });

  return result === confirmLabel;
}

/**
 * 顯示一個可自訂多顆按鈕的確認對話框
 *
 * @param options - 對話框標題、訊息與按鈕設定
 * @returns 使用者實際點擊的按鈕文字（Label）
 *
 * @example
 * const action = await showQuestionDialog({
 *   title: "有新版本",
 *   message: "是否前往下載？",
 *   buttons: [
 *     { label: "取消" },
 *     { label: "前往下載", isDefault: true },
 *   ],
 * });
 *
 * if (action === "前往下載") {
 *   // 執行更新或開啟 Release 頁面
 * }
 */
export async function showQuestionDialog({
  title,
  message,
  buttons,
}: QuestionDialogOptions): Promise<string> {

  const result = await Dialogs.Question({
    Title: title,
    Message: message,
    Buttons: buttons.map((button) => ({
      Label: button.label,
      IsDefault: button.isDefault ?? false,
    })),
  });

  return result;
}
