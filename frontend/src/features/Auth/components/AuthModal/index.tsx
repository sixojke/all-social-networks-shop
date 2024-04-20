import { useState } from "react";

type ContentType = "signIn" | "signUp";

export const AuthModal = () => {
  const [contentType, setContentType] = useState<ContentType>("signIn");
  return <>Auth</>;
};
