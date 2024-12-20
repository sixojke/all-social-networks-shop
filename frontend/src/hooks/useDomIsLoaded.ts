import { useEffect, useState } from "react";

export const useDomIsLoaded = () => {
  const [domLoaded, setDomLoaded] = useState(false);

  useEffect(() => {
    setDomLoaded(true);
  }, []);

  return domLoaded;
};
