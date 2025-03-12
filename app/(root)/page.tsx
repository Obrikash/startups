import Image from "next/image";
import SearchForm from "../../components/SearchForm";
import StartupCard from "@/components/StartupCard";

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ query?: string }>;
}) {
  const query = (await searchParams).query;
  const posts = [
    {
      _createdAt: new Date(),
      views: 55,
      author: { _id: 1, name: "Никита" },
      _id: 1,
      description: "Описание",
      image:
        "https://img.goodfon.ru/wallpaper/big/a/69/kartinka-3d-dikaya-koshka.webp",
      category: "Роботы",
      title: "Мы роботы",
    },
  ];
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
          {posts?.length > 0 ? (
            posts.map((post: StartupTypeCard) => (
              <StartupCard key={post?._id} post={post} />
            ))
          ) : (
            <p className="no-results">No startups found</p>
          )}
        </ul>
      </section>
    </>
  );
}
