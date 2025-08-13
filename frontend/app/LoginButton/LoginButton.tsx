import { Button } from "../components/Button/Button";
import styles from "./LoginButton.module.css";

export const LoginButton = () => {
  return (
    <Button
      img="icons/profile.svg"
      title="Войти"
      className={styles.loginButton}
    />
  );
};
