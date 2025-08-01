import styles from "./page.module.css";
import { PriceCard } from "./components/PriceCard/PriceCard";
import { CategoryCard } from "./components/CategoryCard/CategoryCard";

export default function Home() {
  return (
    <main className={styles.main}>
      {/* <PriceCard
        href="/product"
        img="/beacon.jpg"
        percentDiscount="26"
        title="Сырокопчёный бекон Дымов, нарезка"
        weight="150 г"
        oldPrice="229"
        newPrice="169"
      />
      <PriceCard
        href="/product"
        img="/monterra.jpg"
        percentDiscount="34"
        title="Мороженое Monterra, с кленовым сиропом и грецким орехом, в ведёрке"
        weight="298 г"
        extraText="Снизили цену"
        oldPrice="749"
        newPrice="494"
      /> */}
      <div className={styles.grid}>
        <CategoryCard
          href="/product"
          img="https://cm.samokat.ru/processed/category/11594daf-86aa-4e0f-b623-5b2a2ad66d9e.jpg"
          title="Что на завтрак?"
        />
        <CategoryCard
          href="/product"
          img="https://cm.samokat.ru/processed/category/11594daf-86aa-4e0f-b623-5b2a2ad66d9e.jpg"
          title="Что на завтрак?"
        />
        <CategoryCard
          href="/product"
          img="https://cm.samokat.ru/processed/category/11594daf-86aa-4e0f-b623-5b2a2ad66d9e.jpg"
          title="Что на завтрак?"
        />
        <CategoryCard
          href="/product"
          img="https://cm.samokat.ru/processed/category/11594daf-86aa-4e0f-b623-5b2a2ad66d9e.jpg"
          title="Что на завтрак?"
        />
      </div>
    </main>
  );
}
