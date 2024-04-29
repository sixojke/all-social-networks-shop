import { Button } from "@/shared/components/ui/Buttons/Button";

export const SubcategoryRecord = () => {
  return (
    <div className="w-[60%] h-[1.875vw] flex items-center justify-between bg-white px-[0.625vw] py-[0.625vw] rounded-[0.313vw] border-[#ABABAB] border-[0.016vw]">
      <div className="text-[0.833vw] font-medium">Тестовая категория</div>
      <div className="flex gap-x-[0.417vw]">
        <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
          Удалить
        </Button>
        <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
          Изменить
        </Button>
      </div>
    </div>
  );
};
