import { TABS } from "../../constants/tabs";
import TabList from "@mui/joy/TabList";
import Tab, { tabClasses } from "@mui/joy/Tab";
import Grow from "@mui/material/Grow";
import { useState } from "react";
import classNames from "classnames";

export const Navbar = () => {
  const [activeTab, setActiveTab] = useState<string>("appSettings");
  return (
    <nav className="h-full bg-white pt-8">
      <p className="flex flex-col text-2xl pl-8">
        <span className="font-bold">Admin</span>
        <span>Panel</span>
      </p>
      <TabList
        sx={{
          marginTop: "32px",
          borderTop: "solid 1px #F4F7FE",
          paddingTop: "38px",
          width: "270px",
          rowGap: "10px",
        }}
        disableUnderline
      >
        {Object.keys(TABS).map((key, index) => {
          const targetTabInfo = TABS[key as keyof typeof TABS];
          return (
            <Tab
              indicatorInset
              sx={{
                color: "#D6D6D6",
                [`&.${tabClasses.selected}`]: {
                  bgcolor: "white",
                  color: "#262626",
                  fontWeight: "600",
                  path: {
                    ...(index === 1
                      ? { stroke: "black" }
                      : { fill: "#262626" }),
                  },
                },
                [`&.${tabClasses.root}`]: {
                  height: "27px",
                  padding: "0",
                  paddingLeft: "24px",
                },
                [`&.${tabClasses.root}:after`]: {
                  display: "none",
                },
                [`&.${tabClasses.root}:hover`]: {
                  transitionDuration: "0.3s",
                  bgcolor: "white",
                  color: "#262626",
                  path: {
                    transitionDuration: "0.3s",
                    ...(index === 1
                      ? { stroke: "black" }
                      : { fill: "#262626" }),
                  },
                },
              }}
              key={targetTabInfo.value}
              value={targetTabInfo.value}
            >
              <div
                onClick={() => setActiveTab(targetTabInfo.value)}
                className="flex items-center gap-x-3 justify-between w-full"
              >
                <div className="flex items-center gap-x-4">
                  <targetTabInfo.icon />
                  <span>{targetTabInfo.label}</span>
                </div>
                <Grow
                  in={targetTabInfo.value === activeTab}
                  style={{ transformOrigin: "0 0 0" }}
                  timeout={300}
                >
                  <div
                    className={classNames(
                      "w-[3px] h-[27px] bg-[#262626] rounded-[17px]",
                      {
                        ["invisible"]: targetTabInfo.value !== activeTab,
                      }
                    )}
                  />
                </Grow>
              </div>
            </Tab>
          );
        })}
      </TabList>
    </nav>
  );
};
