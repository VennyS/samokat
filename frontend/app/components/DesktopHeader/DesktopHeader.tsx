import NextLink from "next/link";
import NextImage from "next/image";
import { SearchField } from "../SearchField/SearchField";
import styles from "./DesktopHeader.module.css";
import { LoginButton } from "@/app/LoginButton/LoginButton";
import { ChatButton } from "../ChatButton/ChatButton";

export const DesktopHeader = () => {
  return (
    <header className={styles.header}>
      <NextLink className={styles.logo} scroll={false} href="/">
        <NextImage
          src="https://samokat.ru/images/logo.svg"
          alt="Самокат"
          width={153}
          height={23}
        />
      </NextLink>
      <SearchField
        classname={styles.searchfield}
        placeholder="Искать в Самокатe"
      />
      <div className={styles.rightButtons}>
        <LoginButton />
        <ChatButton />
      </div>
    </header>
  );
};
