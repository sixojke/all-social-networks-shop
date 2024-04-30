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
  Pagination: {
    data: {
      created_at: string;
      description: string;
      referral_code: string;
      total_visitors: number;
    }[];
    limit: number;
    total_items: number;
    total_pages: number;
  };
}
