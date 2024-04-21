import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";
import { Input } from "@/shared/components/ui/Input";
import { Nav } from "../components/Nav";
import { AuthButton } from "@/features/Auth";

export const HeaderLayout = () => {
  return (
    <header className="flex flex-col gap-y-8">
      <section className="flex px-10 gap-x-5 w-full fixed bg-main-white pt-8 z-50 pb-3">
        <Button className="bg-main-dark-blue">PShop</Button>
        <ButtonText className="text-main-dark-blue">Фильтр</ButtonText>
        <Input
          rounded="46px"
          className="!bg-main-light-blue !placeholder-main-light-gray"
          searchIcon
          placeholder="Поиск"
        />
        <AuthButton />
      </section>
      <section className="self-center mt-32">
        <Nav />
      </section>
    </header>
  );
};
