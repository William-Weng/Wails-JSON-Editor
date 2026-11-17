import { visit } from 'jsonc-parser';

export interface JsonErrorLocation {
  line: number;
  column: number;
  offset: number;
}

/**
 * 取得第一個 JSON 語法錯誤的位置。
 * 回傳的行、欄皆從 1 開始。
 */
export function getJsonErrorLocation(
  text: string
): JsonErrorLocation | undefined {
  let location: JsonErrorLocation | undefined;

  visit(
    text,
    {
      onError(_error, offset, _length, startLine, startCharacter) {
        // 只記錄第一個錯誤，避免後續連鎖錯誤干擾。
        if (location) return;

        location = {
          line: startLine + 1,
          column: startCharacter + 1,
          offset
        };
      }
    },
    {
      disallowComments: true,
      allowTrailingComma: false,
      allowEmptyContent: false
    }
  );

  return location;
}

/**
 * 保留原本的錯誤說明，並補上行、欄位置。
 */
export function formatJsonError(
  error: unknown,
  text: string
): string {
  const message =
    error instanceof Error ? error.message : String(error);

  const location = getJsonErrorLocation(text);

  if (!location) {
    return message;
  }

  return `第 ${location.line} 行，第 ${location.column} 欄：${message}`;
}
