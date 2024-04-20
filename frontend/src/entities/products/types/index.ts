export interface IGetAllProductsRequest {
  limit?: number;
  offset?: number;
  category_id?: number;
  subcategory_id?: number;
  is_available?: 0 | 1;
  sort_price?: "asc" | "desc";
  sort_defect?: "asc" | "desc";
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
