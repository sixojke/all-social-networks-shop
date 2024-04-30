const IS_SERVER = typeof window === "undefined";
export const getHostname = (params: { name: string; value: string }[]) => {
  const baseURL = IS_SERVER
    ? process.env.NEXT_PUBLIC_SITE_URL!
    : window.location.origin;
  const url = new URL(baseURL);
  params.forEach((param) => {
    url.searchParams.append(param.name, param.value);
  });
  return url.toString();
};
