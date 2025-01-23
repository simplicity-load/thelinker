import { ENDPOINT, SubmitLinkInput } from "./model";

export async function submitLink(url: string): Promise<Response> {
  const body: SubmitLinkInput = {
    original_url: url,
  };
  const response = await fetch(`${ENDPOINT}/api/v1/shortlink`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  return response.json();
}
