import {useNavigate, useParams} from "react-router";
import {LoginButton} from "@telegram-auth/react";

export const MainPage = () => {
    const navigate = useNavigate();

    const { ref } = useParams<{ ref: string }>();

    return (
        <section className="bg-main-gradient flex h-screen w-full items-center justify-center">
            <div className="min-w-1/3 min-h-1/3 flex flex-col items-center rounded-xl bg-primary p-4">
                <h1 className="text-center text-3xl font-bold text-black dark:text-white">
                    Добро пожаловать в Движ
                </h1>
                <h2 className="mt-10 text-center text-2xl text-black dark:text-white">
                    Вы попали сюда не случайно!
                </h2>
                <h3 className="text-center text-xl text-black dark:text-white">
                    Для получения доступа к чату, авторизуйтесь через Telegram
                </h3>
                <LoginButton botUsername='@zadvizh_assistant_bot' />
                {/*<button*/}
                {/*    onClick={() => navigate(`/auth/continue?ref=${ref}`)}*/}
                {/*    className="mt-10 rounded-lg bg-[#00a8e8] px-4 py-2 text-white duration-200 hover:scale-105"*/}
                {/*>*/}
                {/*    Войти через Telegram*/}
                {/*</button>*/}
            </div>
        </section>
    );
};
