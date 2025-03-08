import Image from "next/image";
import SearchForm from "../../components/SearchForm";

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ query?: string }>;
}) {
  const query = (await searchParams).query;
  return (
    <>
      <section className="pink_container pattern">
        <h1 className="heading">Расскажи про свой Стартап!</h1>
        <p className="sub-heading !max-w-3xl">
          Выложи стартап, голосуй за стартапы и будь замеченным в соревновании
          по стартапам!
        </p>

        <SearchForm query={query} />
      </section>
    </>
  );
}
