-- При виконанні завдання я спочатку написав запит самостійно.
-- Після цього, постфактум, скористався AI для перевірки та можливих оптимізацій.
-- Розумію, що запити можна зробити ефективніше (швидше і синтаксично зрозуміліше), але залишаю свій оригінальний варіант.

-- Номер завдання 2, пункт 2
-- підпункт a
INSERT INTO sensors (room_id, sensor_name, parameter) VALUES
((SELECT room_id FROM rooms WHERE room_name = 'room_A'), 'V1',  'V'),
((SELECT room_id FROM rooms WHERE room_name = 'room_A'), 'R1',  'R'),
((SELECT room_id FROM rooms WHERE room_name = 'room_A'), 'R2',  'R');

-- підпункт b
INSERT INTO sensors (room_id, sensor_name, parameter) VALUES
((SELECT room_id FROM rooms WHERE room_name = 'room_B'), 'V1',  'V'),
((SELECT room_id FROM rooms WHERE room_name = 'room_B'), 'V2',  'V'),
((SELECT room_id FROM rooms WHERE room_name = 'room_B'), 'R1',  'R'),
((SELECT room_id FROM rooms WHERE room_name = 'room_B'), 'R2',  'R'),
((SELECT room_id FROM rooms WHERE room_name = 'room_B'), 'R3',  'R');

-- підпункт с
INSERT INTO measurements (sensor_id, value, ts)
SELECT id, '<value>', '<time>'
FROM sensors WHERE name = 'room_A_V1_V';

INSERT INTO measurements (sensor_id, value, ts)
SELECT id, '<value>', '<time>'
FROM sensors WHERE name = 'room_A_R1_R';


INSERT INTO measurements (sensor_id, value, ts)
SELECT id, '<value>', '<time>'
FROM sensors WHERE name = 'room_A_V1_V';

-- підпункт d
INSERT INTO measurements (sensor_id, value, ts)
SELECT id, '<value>', '<time>'
FROM sensors WHERE name = 'room_B_R1_R';

INSERT INTO measurements (sensor_id, value, ts)
SELECT id, '<value>', '<time>'
FROM sensors WHERE name = 'room_B_R2_R';

-- Номер завдання 2, пункт 3 і далі не мав досвіду побудови таких запитів.
-- Завдяки AI зміг отримати робочий варіантб але, думаю, вас цікавлять мої реальні знання,
-- тому не став вставляти рішення, яке дуже швидко вирішується за допомогою AI, але без допомоги мені вирішити поки доволі складно.
