import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import { useSelector } from 'react-redux';

import { Teacher } from '../../model/Teacher';
import { TeacherCard } from '../TeacherCard/TeacherCard';
import { RootState } from '../../../../app/store/store';

interface TeachersListProps {
    className?: string;
}

export const TeachersList = (props: TeachersListProps) => {
    const { className } = props;

    const navigate = useNavigate();
    const userToken = useSelector((state: RootState) => state.user.accessToken);

    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [teachers, setTeachers] = useState<Teacher[]>([]);

    useEffect(() => {
        const fetchTeachers = async () => {
            const result = await fetch('https://zadvizh.tech/api/v1/teachers', {
                headers: {
                    Authorization: userToken || '',
                },
            });

            if (result.ok) {
                setTeachers(await result.json());
            }
        };

        try {
            setIsLoading(true);
            fetchTeachers();
        } catch (e) {
            console.log(e);
            alert('Что-то сломалось ');
            navigate('/');
        } finally {
            setIsLoading(false);
        }
    }, []);

    if (isLoading) {
        return (
            <div>
                <h1 className="text-xl italic opacity-40 lg:text-3xl">Загрузка данных...</h1>
            </div>
        );
    }

    if (!teachers?.length) {
        return (
            <div>
                <h1 className="text-xl italic opacity-40 lg:text-3xl">Ничего не нашли...(</h1>
            </div>
        );
    }

    return (
        <div className={'flex w-full flex-col gap-3 ' + className}>
            {teachers.map((t) => (
                <TeacherCard teacher={t} key={t.id} />
            ))}
        </div>
    );
};
