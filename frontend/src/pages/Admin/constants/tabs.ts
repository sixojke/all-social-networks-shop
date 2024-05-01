import HomeIcon from "@/assets/adminIcons/home.svg";
import CartIcon from "@/assets/adminIcons/cart.svg";
import CategoriesIcon from "@/assets/adminIcons/bolcks.svg";
import ProfileIcon from "@/assets/adminIcons/user.svg";
import InfoIcon from "@/assets/adminIcons/info.svg";
import ReferalIcon from "@/assets/adminIcons/referals.svg";
import AdminLogsIcon from "@/assets/adminIcons/adminActions.svg";

export const TABS = {
  appSettings: {
    icon: HomeIcon,
    value: "appSettings",
    label: "Настройки сайта",
  },
  referrals: {
    icon: ReferalIcon,
    value: "referrals",
    label: "Реферальные ссылки",
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
  adminLogs: {
    value: "adminLogs",
    label: "Действия админов",
    icon: AdminLogsIcon,
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
