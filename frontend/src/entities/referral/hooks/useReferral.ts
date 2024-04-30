import { useSearchParams } from "next/navigation";
import { useAddVisitorMutation } from "../service";
import { useEffect } from "react";

export const useReferral = () => {
  const [addVisitor] = useAddVisitorMutation();
  const searchParams = useSearchParams();
  useEffect(() => {
    const code = searchParams?.get("code");
    const localCode = localStorage.getItem("code");
    if (!code) return;
    if (code === localCode) return;
    localStorage.setItem("code", code);
    addVisitor({ referral_code: code });
  }, [searchParams]);
};
