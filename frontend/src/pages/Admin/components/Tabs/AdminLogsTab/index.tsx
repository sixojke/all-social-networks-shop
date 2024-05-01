import TabPanel from "@mui/joy/TabPanel";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { QueryHandler } from "@/widgets/QueryHandler";
import { Pagination } from "@/shared/components/ui/Pagination";
import { useState } from "react";
import { useGetAdminLogsQuery } from "@/entities/logs";
import { LogRecord } from "./components/LogRecord";

export const AdminLogTab = () => {
  const [page, setPage] = useState<number>(1);
  const onPageChange = (page: number) => {
    setPage(page + 1);
  };
  const { data, isError, isLoading } = useGetAdminLogsQuery({
    limit: 10,
    page,
  });

  return (
    <TabPanel
      sx={{
        [`&.MuiTabPanel-root`]: {
          padding: 0,
        },
      }}
      value="adminLogs"
    >
      <div className="flex flex-col justify-between h-[100vh] pb-[1.563vw]">
        <div>
          <div className="flex items-center justify-between">
            <p className="text-[1.302vw] font-bold">Все ссылки</p>
          </div>
          <QueryHandler
            isError={isError}
            isLoading={isLoading}
            errorLabel="Произошла ошибка при загрузке логов"
          />
          {!data?.Pagination.data && !isError && !isLoading && (
            <div className="w-full flex justify-center flex-col items-center gap-y-9 mt-24 text-3xl font-semibold text-main-dark-green">
              Действия на найдены
            </div>
          )}
          <div className="mt-[0.625vw] flex flex-col gap-y-[0.625vw]">
            {data?.Pagination?.data?.map((log) => {
              return (
                <LogRecord
                  key={log.created_at}
                  actionDate={log.created_at}
                  adminName={log.username}
                  adminAction={log.message}
                />
              );
            })}
          </div>
        </div>
        {!!data?.Pagination && !!data?.Pagination.total_pages && (
          <div className="mt-[40px] mb-[4.167vw] flex justify-center">
            <Pagination
              countPage={data.Pagination.total_pages}
              currentPage={page}
              onChange={onPageChange}
            />
          </div>
        )}
      </div>
    </TabPanel>
  );
};
