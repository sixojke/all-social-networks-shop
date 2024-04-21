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
  min_hold_time: number;
  category_id: number;
}
