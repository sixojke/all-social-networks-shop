export interface IGetAllCategoriesResponse {
  data: ICategory[] | null;
  count: number;
}

export interface ICategory {
  id: number;
  name: string;
  img_path: string;
}
