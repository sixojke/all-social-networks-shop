import { useTestQuery } from "./service";

export const TestPage = () => {
  const { data } = useTestQuery({});
  console.log(data);
  return <>{data ? "Авторизован" : "Не авторизовн"}</>;
};
