import { Button } from "../Button/Button";
import styles from "./ChatButton.module.css";

export const ChatButton = () => {
  return <Button img="icons/chat.svg" className={styles.chatButton} />;
};
