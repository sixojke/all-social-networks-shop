import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";
import { Input } from "@/shared/components/ui/Input";
import { Nav } from "../components/Nav";
import { AuthButton } from "@/features/Auth";

export const HeaderLayout = () => {
  return (
    <header className="flex flex-col gap-y-8">
      <section className="flex mt-8 px-10 gap-x-5">
        <Button className="bg-main-dark-blue">PShop</Button>
        <ButtonText className="text-main-dark-blue">Фильтр</ButtonText>
        <Input placeholder="Поиск" />
        <AuthButton />
      </section>
      <section className="self-center">
        <Nav />
      </section>
    </header>
  );
};
