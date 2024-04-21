export interface IGetAllProductsRequest {
  limit?: number;
  page?: number;
  category_id?: number;
  subcategory_id?: number;
  is_available?: 0 | 1;
  sort_price?: "asc" | "desc";
  sort_defect?: "asc" | "desc";
}

export interface IGetAllProductsResponse {
  Pagination: {
    data: Product[] | null;
    total_pages: number;
    total_items: number;
    limit: number;
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
