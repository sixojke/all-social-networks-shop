import HomeIcon from "@/assets/adminIcons/home.svg";
import CartIcon from "@/assets/adminIcons/cart.svg";
import CategoriesIcon from "@/assets/adminIcons/bolcks.svg";
import ProfileIcon from "@/assets/adminIcons/user.svg";
import InfoIcon from "@/assets/adminIcons/info.svg";

export const TABS = {
  appSettings: {
    icon: HomeIcon,
    value: "appSettings",
    label: "Настройки сайта",
  },
  products: {
    icon: CartIcon,
    value: "products",
    label: "Товары",
  },
  categories: {
    value: "categories",
    label: "Категории",
    icon: CategoriesIcon,
  },
  profile: {
    value: "profile",
    label: "Профиль",
    icon: ProfileIcon,
  },
  arbitration: {
    value: "arbitration",
    label: "Арбитраж",
    icon: InfoIcon,
  },
};
