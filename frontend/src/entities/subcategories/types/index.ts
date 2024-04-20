export interface IGetAllSubcategoriesRequest {
  category_id: number;
}

export interface IGetAllSubcategoriesResponse {
  data: ISubcategory[] | null;
  count: number;
}

export interface ISubcategory {
  id: number;
  name: string;
  hold_time: number;
  category_id: number;
}
