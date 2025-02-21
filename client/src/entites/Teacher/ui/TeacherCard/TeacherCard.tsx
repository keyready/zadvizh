import { Button } from '@heroui/button';
import { RiDislikeFill, RiHeart2Fill, RiSendPlane2Line } from '@remixicon/react';
import { useCallback, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { Textarea } from '@heroui/input';

import { Teacher } from '../../model/Teacher.ts';
import { RootState } from '../../../../app/store/store.ts';

interface TeacherCardProps {
    className?: string;
    teacher: Teacher;
}

export const TeacherCard = (props: TeacherCardProps) => {
    const { teacher, className } = props;

    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [isCommentsVisible, setIsCommentsVisible] = useState<boolean>(false);

    const userId = useSelector((state: RootState) => state.user.id);
    const userToken = useSelector((state: RootState) => state.user.accessToken);

    const [comment, setComment] = useState<string>('');

    const hasAuthorPromoted = useMemo(() => {
        return {
            likes: teacher.likes.authors.includes(userId),
            dislikes: teacher.dislikes.authors.includes(userId),
        };
    }, []);

    const handleDislikePress = useCallback(async () => {
        setIsLoading(true);

        try {
            await fetch('https://zadvizh.tech/api/v1/teacher/likes', {
                method: 'post',
                body: JSON.stringify({ teacherId: teacher.id, action: 'like' }),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': userToken || '',
                },
            });
        } catch (e) {
            alert('Что-то сломалось ' + e);
        } finally {
            setIsLoading(false);
        }
    }, [userToken, teacher?.id]);

    const handleLikePress = useCallback(async () => {
        setIsLoading(true);

        try {
            await fetch('https://zadvizh.tech/api/v1/teacher/likes', {
                method: 'post',
                body: JSON.stringify({ teacherId: teacher.id, action: 'dislike' }),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': userToken || '',
                },
            });
        } catch (e) {
            alert('Что-то сломалось ' + e);
        } finally {
            setIsLoading(false);
        }
    }, [userToken, teacher?.id]);

    const handleSendComment = useCallback(async () => {
        setIsLoading(true);

        try {
            await fetch('https://zadvizh.tech/api/v1/teachers/addComment', {
                method: 'post',
                body: JSON.stringify({ teacherId: teacher.id, content: comment }),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': userToken || '',
                },
            });
        } catch (e) {
            alert('Что-то сломалось ' + e);
        } finally {
            setIsLoading(false);
        }
    }, [comment, userToken, teacher?.id]);

    return (
        <div
            className={
                'flex w-full flex-col gap-2 rounded-md border-1 border-success p-3 backdrop-blur-lg ' +
                className
            }
        >
            <h1 className="text-2xl font-bold">
                {teacher.lastname} {teacher.firstname.slice(0, 1)}. {teacher.middlename.slice(0, 1)}
                .
            </h1>
            <div className="flex gap-2">
                <Button
                    isLoading={isLoading}
                    isDisabled={hasAuthorPromoted.dislikes}
                    onPress={handleLikePress}
                    variant="bordered"
                    color={hasAuthorPromoted.likes ? 'danger' : 'primary'}
                >
                    <RiHeart2Fill
                        className={hasAuthorPromoted.likes ? 'text-danger' : 'text-primary'}
                    />
                    {teacher.likes.value}
                </Button>

                <Button
                    isLoading={isLoading}
                    isDisabled={hasAuthorPromoted.likes}
                    onPress={handleDislikePress}
                    variant="bordered"
                    color={hasAuthorPromoted.dislikes ? 'danger' : 'primary'}
                >
                    <RiDislikeFill
                        className={hasAuthorPromoted.dislikes ? 'text-danger' : 'text-primary'}
                    />
                    {teacher.dislikes.value}
                </Button>
            </div>

            {hasAuthorPromoted.dislikes ||
                (hasAuthorPromoted.likes && (
                    <div className="mt-4 flex flex-col gap-2">
                        <h2>Вы можете написать комментарий</h2>
                        <Textarea
                            value={comment}
                            onValueChange={setComment}
                            endContent={
                                <Button
                                    onPress={handleSendComment}
                                    color="success"
                                    isDisabled={!comment}
                                    variant="bordered"
                                >
                                    <RiSendPlane2Line />
                                </Button>
                            }
                            classNames={{
                                inputWrapper: 'h-auto',
                            }}
                            maxRows={10}
                            variant="bordered"
                            placeholder="Ваш комментарий для преподавателя"
                        />
                    </div>
                ))}

            {teacher?.comments?.length ? (
                <div className="mt-5">
                    <button className="mb-3" onClick={() => setIsCommentsVisible((ps) => !ps)}>
                        {isCommentsVisible ? 'Скрыть' : 'Показать'} комментарии
                    </button>
                    {isCommentsVisible &&
                        teacher?.comments.map((com) => (
                            <div key={com.id} className="rounded-md border-1 border-white p-4">
                                <p className="opacity-40">
                                    {new Date(com.createdAt).toLocaleDateString()}
                                </p>
                                <h1>{com.content}</h1>
                            </div>
                        ))}
                </div>
            ) : null}
        </div>
    );
};
