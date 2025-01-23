import { ENDPOINT } from "./model";

export async function getLinks(): Promise<Response> {
  const response = await fetch(`${ENDPOINT}/api/v1/shortlinks`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const { data } = await response.json();
  return data;
}
