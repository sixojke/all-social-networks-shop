import { Button } from "@/shared/components/ui/Buttons/Button";

export const CategoryRecord = () => {
  return (
    <div className="w-full h-9 flex items-center justify-between bg-white px-3 py-3 rounded-[6px]">
      <div>Test</div>
      <div>
        <Button>Удалить</Button>
        <Button>Изменить</Button>
      </div>
    </div>
  );
};
