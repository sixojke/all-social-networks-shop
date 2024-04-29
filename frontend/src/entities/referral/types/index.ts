export interface ICreateLinkRequest {
  description: string;
}

export interface ICreateLinkResponse {
  link: string;
}

export interface IDeleteLinkRequest {
  code: string;
}

export interface IAddVisitorRequest {
  referral_code: string;
}

export interface IGetStatsRequest {
  limit: number;
  page: number;
}

export interface IGetStatsResponse {
  pagination: {
    data: "string";
    limit: 0;
    total_items: 0;
    total_pages: 0;
  };
}
