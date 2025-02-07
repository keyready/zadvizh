import {useNavigate, useParams} from "react-router";

export const MainPage = () => {
    const navigate = useNavigate()

    const {ref} = useParams<{ ref: string }>();

    return (
        <section className="flex items-center justify-center w-full h-screen bg-main-gradient">
            <div className="flex flex-col items-center p-4 w-1/3 h-1/3 bg-primary rounded-xl">
                <h1 className="text-center font-bold text-black dark:text-white text-3xl">Добро пожаловать в
                    Движ</h1>
                <h2 className="mt-10 text-center text-black dark:text-white text-2xl">Вы попали сюда не
                    случайно!</h2>
                <h3 className="text-center text-black dark:text-white text-xl">Для получения доступа к чату,
                    авторизуйтесь через Telegram</h3>
                <button
                    onClick={() => navigate(`/auth/continue?ref=${ref}`)}
                    className="mt-10 bg-[#00a8e8] text-white hover:scale-105 duration-200 rounded-lg py-2 px-4">Войти
                    через Telegram
                </button>
            </div>
        </section>
    );
};
