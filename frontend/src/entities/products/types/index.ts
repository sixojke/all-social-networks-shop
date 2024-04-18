export interface IGetAllProductsRequest {
  limit?: number;
  offset?: number;
}

export interface IGetAllProductsResponse {
  Pagination: {
    data: [
      {
        id: number;
        name: string;
        price: number;
        quantity: number;
        quantity_sales: number;
        description: string;
        img: string;
        uploaded_at: string;
        category_id: 2;
      }
    ];
    total: number;
    limit: number;
    offset: number;
  };
}
