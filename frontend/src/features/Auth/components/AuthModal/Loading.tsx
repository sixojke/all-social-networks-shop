import { Loader } from "@/shared/components/ui/Loader";

export const Loading = () => {
  return (
    <div className="w-full mt-32 flex flex-col gap-y-5 items-center justify-center">
      <Loader />
    </div>
  );
};
