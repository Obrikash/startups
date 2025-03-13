import Image from "next/image";
import SearchForm from "../../components/SearchForm";
import StartupCard from "@/components/StartupCard";

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ title?: string }>;
}) {
  const query = (await searchParams).title;
  const response = await fetch(
    `http://localhost:4000/api/startups${query ? `?title=${query}` : ""}`
  );
  const posts = await response.json();
  console.log(JSON.stringify(posts));
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
      <section className="section_container">
        <p className="text-30-semibold">
          {query ? `Результаты поиска для "${query}"` : "Все стартапы"}
        </p>
        <ul className="mt-7 card_grid">
          {posts?.metadata?.total_records > 0 ? (
            posts?.startups.map((post: StartupTypeCard) => (
              <StartupCard key={post?.id} post={post} />
            ))
          ) : (
            <p className="no-results">No startups found</p>
          )}
        </ul>
      </section>
    </>
  );
}
