import unittest


class CircularQueue:
    def __init__(self, size):
        self.element = [None] * size
        self.size = size
        self.front = 0
        self.back = -1

    def enqueue(self, element):
        self.back += 1
        self.element[self.back % self.size] = element

    def dequeue(self):
        if self.is_empty():
            raise ValueError("No elements in the queue")
        value = self.get_front()
        self.element[self.front % self.size] = None
        self.front += 1
        return value

    def get_front(self):
        if self.is_empty():
            raise ValueError("No elements in the queue")
        return self.element[self.front % self.size]

    def is_empty(self):
        return all(x is None for x in self.element)

    def get_length(self):
        return sum(x is not None for x in self.element)


class TestCircularQueue(unittest.TestCase):
    def setUp(self):
        self.queue = CircularQueue(5)

    def test_enqueue(self):
        self.queue.enqueue(1)
        self.assertEqual(self.queue.get_front(), 1)
        self.assertEqual(self.queue.get_length(), 1)

    def test_dequeue(self):
        self.queue.enqueue(1)
        self.queue.enqueue(2)
        self.assertEqual(self.queue.dequeue(), 1)
        self.assertEqual(self.queue.get_length(), 1)

    def test_get_front(self):
        self.queue.enqueue(1)
        self.assertEqual(self.queue.get_front(), 1)

    def test_is_empty(self):
        self.assertTrue(self.queue.is_empty())
        self.queue.enqueue(1)
        self.assertFalse(self.queue.is_empty())

    def test_get_length(self):
        self.assertEqual(self.queue.get_length(), 0)
        self.queue.enqueue(1)
        self.queue.enqueue(2)
        self.assertEqual(self.queue.get_length(), 2)


if __name__ == '__main__':
    unittest.main()
