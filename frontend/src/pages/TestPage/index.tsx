import { useTestQuery } from "./service";

export const TestPage = () => {
  const { data } = useTestQuery({});
  return <>{data ? "Авторизован" : "Не авторизовн"}</>;
};
