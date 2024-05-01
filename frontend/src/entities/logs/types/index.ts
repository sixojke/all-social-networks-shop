export interface IGetAdminLogsRequest {
  limit?: number;
  page?: number;
}

export interface IGetAdminLogsResponse {
  Pagination: {
    data: {
      created_at: string;
      message: string;
      username: string;
    }[];
    limit: number;
    total_items: number;
    total_pages: number;
  };
}
