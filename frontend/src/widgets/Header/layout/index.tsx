import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";
import { Input } from "@/shared/components/ui/Input";
import { Nav } from "../components/Nav";
import { AuthButton, useLogout } from "@/features/Auth";
import { useGetActiveUserQuery } from "@/entities/user";

export const HeaderLayout = () => {
  const { data: activeUser } = useGetActiveUserQuery();
  const logout = useLogout();
  return (
    <header className="flex flex-col gap-y-8">
      <section className="flex px-10 gap-x-5 w-full fixed bg-[#FFFFFF] pt-8 z-50 pb-3 items-center">
        <Button className="bg-main-dark-green">PShop</Button>
        <ButtonText className="text-main-dark-green">Фильтр</ButtonText>
        <Input
          wrapperClassname="!border-main-black !bg-main-light-gray !border-solid !border-opacity-20 !border-[0.5px]"
          className="!bg-main-light-gray !text-main-black !bg-opacity-60  !placeholder-main-dark-gray"
          searchIcon
          placeholder="Поиск"
        />
        {activeUser ? (
          <>
            {activeUser.username} <Button onClick={logout}>Выйти</Button>
          </>
        ) : (
          <AuthButton />
        )}
      </section>
      <section className="self-center mt-32">
        <Nav />
      </section>
    </header>
  );
};
