import matplotlib.pyplot as plt
import tensorflow as tf

""" Week 1 example
model = tf.keras.Sequential([keras.layers.Dense(units=100, input_shape=[1])])
model.compile(optimizer='sgd', loss='mean_squared_error')

xs = np.array([-1.0,  0.0, 1.0, 2.0, 3.0, 4.0], dtype=float)
ys = np.array([-3.0, -1.0, 1.0, 3.0, 5.0, 7.0], dtype=float)

model.fit(xs, ys, epochs=1000)

print(model.predict([100]))
"""

""" Week 1 task
def house_model(y_new):
    model = tf.keras.Sequential([keras.layers.Dense(units=1, input_shape=[1])])
    model.compile(optimizer='sgd', loss='mean_squared_error')
    xs = np.array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0], dtype=float)
    ys = np.array([1.0, 1.5, 2.0, 2.5, 3.0, 3.5], dtype=float)
    model.fit(xs, ys, epochs=1000, use_multiprocessing=True)
    return model.predict(y_new)[0]

prediction = house_model([7.0])
print(prediction)
"""

""" 
Week 2 example 
https://colab.research.google.com/github/lmoroney/dlaicourse/blob/master/Course%201%20-%20Part%204%20-%20Lesson%202%20-%20Notebook.ipynb
"""


class myCallback(tf.keras.callbacks.Callback):
    def on_epoch_end(self, epoch, logs={}):
        if (logs.get('loss') < 0.4):
            print("\nReached 60% accuracy so cancelling training!")
            self.model.stop_training = True


mnist = tf.keras.datasets.fashion_mnist
(training_images, training_labels), (test_images, test_labels) = mnist.load_data()
plt.imshow(training_images[0])
print(training_labels[0])
print(training_images[0])

# normilise (0..255 -> 0..1)
training_images = training_images / 255.0
test_images = test_images / 255.0

model = tf.keras.models.Sequential([  # That defines a SEQUENCE of layers in the neural network
    tf.keras.layers.Flatten(),  # Flatten just takes that square and turns it into a 1 dimensional set.
    # Adds a layer of neurons with activation function Relu - effectively mea   ns "If X>0 return X, else return 0" -- so what it does it only passes values 0 or greater to the next layer in the network.
    tf.keras.layers.Dense(1024, activation=tf.nn.relu),
    # Adds a layer with activation function Softman - takes a set of values, and effectively picks the biggest one,
    # so, for example, if the output of the last layer looks like [0.1, 0.1, 0.05, 0.1, 9.5, 0.1, 0.05, 0.05, 0.05],
    # it saves you from fishing through it looking for the biggest value, and turns it into [0,0,0,0,1,0,0,0,0] -- The goal is to save a lot of coding!
    tf.keras.layers.Dense(10, activation=tf.nn.softmax)
])

model.compile(optimizer=tf.optimizers.Adam(), loss="sparse_categorical_crossentropy", metrics=['accuracy'])
model.fit(training_images, training_labels, epochs=10, callbacks=[myCallback()])
model.evaluate(test_images, test_labels)

classifications = model.predict(test_images)
print(classifications[0])
print(test_labels[0])
