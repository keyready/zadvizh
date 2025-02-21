import { Divider } from '@heroui/react';

import { TeachersList } from '../../../entites/Teacher';

export const RatingPage = () => {
    return (
        <section className="bg-main-gradient relative flex h-screen w-full flex-col items-start justify-start overflow-y-auto px-3 py-20 lg:px-[15vw]">
            <div className="flex flex-col gap-3">
                <h1 className="text-xl font-bold lg:text-3xl">
                    Рейтинг проффесорско-преподавательского состава
                </h1>
                <p className="text-md lg:text-2xl">
                    Здесь Вы можете оценить работу каждого из преподавателей кафедры и написать
                    отзыв о его работе
                </p>
                <p className="text-md italic opacity-40 lg:text-2xl">
                    <b>ВАЖНО</b>: мы не храним в открытом виде информацию о том, кто поставил какие
                    оценки. ID пользователей хранится в зашифрованном в одностороннем порядке виде и
                    нужно только для того, чтобы оценки не будлировались
                </p>
            </div>

            <Divider className="my-5 w-full" />

            <TeachersList />
        </section>
    );
};
