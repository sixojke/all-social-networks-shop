import { useGetActiveUserQuery } from "@/entities/user";
import { PermissionTabs } from "../types";

export const useGetPermissionTabs = (tabs: PermissionTabs) => {
  const { data: activeUser } = useGetActiveUserQuery();
  let permissionTabs: {
    [key: string]: {
      icon: any;
      value: string;
      label: string;
    };
  } = {};
  switch (activeUser?.role) {
    case "admin":
      permissionTabs = tabs;
      break;
    default:
      permissionTabs["profile"] = tabs.profile;
  }
  return permissionTabs;
};
