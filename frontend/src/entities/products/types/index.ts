export interface IGetAllProductsRequest {
  limit?: number;
  offset?: number;
}

export interface IGetAllProductsResponse {
  Pagination: {
    data: Product[] | null;
    total: number;
    limit: number;
    offset: number;
  };
}

export type Product = {
  id: number;
  name: string;
  price: number;
  quantity: number;
  quantity_sales: number;
  description: string;
  img: string;
  uploaded_at: string;
  category_id: 2;
};