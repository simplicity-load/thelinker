export interface Response {
  data: any;
  error?: string | null;
}

export const ENDPOINT = "http://localhost:8800";

export interface SubmitLinkInput {
  original_url: string;
}

export interface Shortlink {
  original_url: string;
  short_url: string;
  date: string;
}
