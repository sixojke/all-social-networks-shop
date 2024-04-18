export interface IOptionsSelect {
  id: number;
  name: string;
}

export interface IOption extends NonNullable<unknown> {
  id: number | string;
  name: string;
}

export interface IOptionsGroupSelect {
  label: string;
  options: IOption[];
}
