import pandas as pd

# Prices in NT Dollars for iPhone 15 Pro models
prices = {
    128: 36900,
    256: 40400,
    512: 47400,
    1024: 54400,  # 1TB is 1024GB
}

# Initialize a DataFrame from the prices dictionary
df = pd.DataFrame(list(prices.items()), columns=["GB", "Price"])
# Calculate and add a column for price per GB
df["Price Per GB"] = df["Price"] / df["GB"]

# Initialize a DataFrame for price differences
diff = pd.DataFrame()

# Calculate additional price and GB for each upgrade
diff["Add Price"] = df["Price"].diff().shift(-1)
diff["Add GB"] = df["GB"].diff().shift(-1)

# Drop the last row from both dataframes as it does not have a valid comparison
df = df[:-1]
diff = diff[:-1]

# Calculate Price Difference Per Additional GB and round to 2 decimal places
diff["Price Diff Per Add GB"] = (diff["Add Price"] / diff["Add GB"]).round(2)

# Convert Add GB to integer
diff["Add GB"] = diff["Add GB"].astype(int)

# Adding 'spec' column for storage upgrade specification
diff["GB Diff"] = df["GB"].astype(str) + " to " + (df["GB"] + diff["Add GB"]).astype(str)

# Reordering columns for clarity
diff = diff[["GB Diff", "Add Price", "Add GB", "Price Diff Per Add GB"]]

# Output the DataFrames
print("Data DataFrame:")
print(df)
print("\nPrice Difference DataFrame:")
print(diff)
