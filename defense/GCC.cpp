#include <iostream>
#include <cmath>
#include <vector>
#include <random>

const double PI = 3.14159265358979323846;

// Gaussian function
double gaussian(double x, double mean, double stddev) {
    return (1.0 / (stddev * sqrt(2 * PI))) * exp(-0.5 * pow((x - mean) / stddev, 2));
}

// Key generation using Gaussian distribution
std::vector<double> generateKey(int size, double mean, double stddev) {
    std::vector<double> key(size);
    std::default_random_engine generator;
    std::normal_distribution<double> distribution(mean, stddev);

    for (int i = 0; i < size; ++i) {
        key[i] = distribution(generator);
    }
    return key;
}

// Encrypt message
std::vector<double> encrypt(const std::string& message, const std::vector<double>& key) {
    std::vector<double> encrypted;
    for (size_t i = 0; i < message.size(); ++i) {
        encrypted.push_back(static_cast<double>(message[i]) + key[i % key.size()]);
    }
    return encrypted;
}

// Decrypt message
std::string decrypt(const std::vector<double>& encrypted, const std::vector<double>& key) {
    std::string decrypted;
    for (size_t i = 0; i < encrypted.size(); ++i) {
        decrypted.push_back(static_cast<char>(encrypted[i] - key[i % key.size()]));
    }
    return decrypted;
}

int main() {
    std::string message = "Hello, World!";
    double mean = 0.0;
    double stddev = 1.0;

    // Key generation
    std::vector<double> key = generateKey(message.size(), mean, stddev);

    // Encryption
    std::vector<double> encrypted = encrypt(message, key);

    // Display encrypted data
    std::cout << "Encrypted: ";
    for (double val : encrypted) {
        std::cout << val << " ";
    }
    std::cout << std::endl;

    // Decryption
    std::string decrypted = decrypt(encrypted, key);
    std::cout << "Decrypted: " << decrypted << std::endl;

    return 0;
}

