import { TABS } from "../../constants/tabs";
import TabList from "@mui/joy/TabList";
import Tab, { tabClasses } from "@mui/joy/Tab";

export const Navbar = () => {
  return (
    <nav className="h-full bg-white pt-[1.667vw]">
      <p className=" text-[0.99vw] leading-[1.042vw] pl-[1.667vw]">
        <span className="font-bold">Admin</span>
        <br />
        <span>Panel</span>
      </p>
      <TabList
        sx={{
          marginTop: "1.146vw",
          borderTop: "solid 1px #F4F7FE",
          paddingTop: "1.563vw",
          width: "11.302vw",
          rowGap: "0.521vw",
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
                  ["line"]: {
                    visibility: "visible !important",
                  },
                },
                [`&.${tabClasses.root}`]: {
                  height: "1.406vw",
                  padding: "0",
                  paddingLeft: "1.25vw",
                },
                [`&.${tabClasses.root}:after`]: {
                  width: "0.156vw",
                  height: "1.406vw",
                  borderRadius: "0.938vw",
                  transition: "0.5s",
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
              <div className="flex items-center gap-x-3 justify-between w-full">
                <div className="flex items-center gap-x-[12px]">
                  <targetTabInfo.icon/>
                  <span className="text-[0.625vw]">{targetTabInfo.label}</span>
                </div>
              </div>
            </Tab>
          );
        })}
      </TabList>
    </nav>
  );
};
